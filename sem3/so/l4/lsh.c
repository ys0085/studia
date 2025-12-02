#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <fcntl.h>
#include <errno.h>


void child_handler(int sig){
    pid_t pid;
    int status;

    while(waitpid(-1, &status, WNOHANG) > 0);
}
    
char **parse_line(char *line, int *background_flag) {
    char **args = malloc(100 * sizeof(char *));
    char *token;
    int position = 0;

    *background_flag = 0;

    token = strtok(line, " \n\t");
    while (token != NULL) {
        args[position++] = token;
        token = strtok(NULL, " \n\t");
    }
    if (position > 0 && strcmp(args[position - 1], "&") == 0) {
        *background_flag = 1;
        args[position - 1] = NULL; 
    }
    args[position] = NULL;
    return args;
}

char *read_line() {
    char *line = NULL;
    size_t bufsize = 0;
    if (getline(&line, &bufsize, stdin) == -1) {
        if (feof(stdin)) {
            exit(EXIT_SUCCESS); 
        } else {
            fprintf(stderr, "getline error\n");
            exit(EXIT_FAILURE);
        }
    }
    return line;
}

void handle_redirections(char **args) {
    for (int i = 0; args[i] != NULL; i++) {
        if (strcmp(args[i], ">") == 0) {
            int fd = open(args[i + 1], O_WRONLY | O_CREAT | O_TRUNC, 0644);
            if (fd == -1) {
                perror("error");
                exit(EXIT_FAILURE);
            }
            dup2(fd, STDOUT_FILENO);
            close(fd);
            args[i] = NULL;
        } else if (strcmp(args[i], "2>") == 0) {
            int fd = open(args[i + 1], O_WRONLY | O_CREAT | O_TRUNC, 0644);
            if (fd == -1) {
                perror("error");
                exit(EXIT_FAILURE);
            }
            dup2(fd, STDERR_FILENO);
            close(fd);
            args[i] = NULL;
        } else if (strcmp(args[i], "<") == 0) {
            int fd = open(args[i + 1], O_RDONLY);
            if (fd == -1) {
                perror("error");
                exit(EXIT_FAILURE);
            }
            dup2(fd, STDIN_FILENO);
            close(fd);
            args[i] = NULL;
        }
    }
}

int lsh_exit(char **args) {
    return 0;
}

int lsh_cd(char **args){
    if (args[1] == NULL) {
        fprintf(stderr, "expected argument\n");
    } else {
        if (chdir(args[1]) != 0) {
            fprintf(stderr, "wrong argument\n");
        }
    }
    return 1;
}

int lsh_execute(char **args, int background_flag) {
    pid_t cpid;
    int status;
    cpid = fork();

    if (cpid == 0) {
        if (execvp(args[0], args) < 0){
            printf("command not found: %s\n", args[0]);
            exit(EXIT_FAILURE);
        }

    } 
    else if (cpid < 0) printf("error forking\n");
    else {
        if(!background_flag) waitpid(cpid, NULL, 0);
        else printf("process running in background with PID %d\n", cpid);
    }

    return 1;
}

void lsh_pipe(char **cmd1, char **cmd2, int background_flag) {
    int pipefd[2];
    pid_t pid1, pid2;

    if (pipe(pipefd) == -1) {
        perror("pipe");
        exit(EXIT_FAILURE);
    }

    pid1 = fork();
    if (pid1 == 0) {
        // Proces pierwszy: zapisuje do potoku
        close(pipefd[0]);          // Zamykamy czytanie
        dup2(pipefd[1], STDOUT_FILENO); // Przekierowanie wyjścia do potoku
        close(pipefd[1]);
        execute_single_command(cmd1);
    } else if (pid1 < 0) {
        perror("fork");
        exit(EXIT_FAILURE);
    }

    pid2 = fork();
    if (pid2 == 0) {
        // Proces drugi: czyta z potoku
        close(pipefd[1]);          // Zamykamy zapis
        dup2(pipefd[0], STDIN_FILENO); // Przekierowanie wejścia z potoku
        close(pipefd[0]);
        execute_single_command(cmd2);
    } else if (pid2 < 0) {
        perror("fork");
        exit(EXIT_FAILURE);
    }

    // Proces macierzysty: zamyka potok i czeka na oba procesy
    close(pipefd[0]);
    close(pipefd[1]);
    waitpid(pid1, NULL, 0);
    waitpid(pid2, NULL, 0);
}

void parse_pipe(char *line){
    char *line1 = strtok(line, "|");
    char *line2 = strtok(NULL, "|");

    int background_flag = 0;

    char **cmd1 = parse_line(line1, &background_flag);
    char **cmd2 = parse_line(line2, &background_flag);

    lsh_pipe(cmd1, cmd2, background_flag);
}

void loop(){
    char *line;
    char **args;
    char **args2;
    
    int status = 1;

    int background_flag = 0;

    char *cmd1;
    char *cmd2;

    struct sigaction sa;
    sa.sa_flags = SA_RESTART;
    sa.sa_handler = child_handler;
    sigaction(SIGCHLD, &sa, NULL);

    do {
        printf("lsh> ");
        line = read_line();

        if (strchr(line, '|')) parse_pipe(line);

        background_flag = 0;
        args = parse_line(line, &background_flag);
        if (args == NULL || args[0] == NULL) {
            free(line);
            free(args);
            continue;
        }

        if(strcmp(args[0], "cd") == 0) status = lsh_cd(args);
        else if(strcmp(args[0], "exit") == 0) status = lsh_exit(args);
        else status = lsh_execute(args, background_flag);

        free(line);
        free(args);

    } while (status);
}

int main(int argc, char **argv){
    loop();
}