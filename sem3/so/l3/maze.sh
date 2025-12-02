maze=()
wall_char="█";

maze_width=$(tput cols)
maze_height=$(tput lines)

maze_width=$(((maze_width - 2) * 2 / 2 + 1)) # Ensures the maze's dimensions are odd numbers. (it looks better)
maze_height=$(((maze_height - 2) * 2 / 2 + 1))

# Prim's algorithm

# Step 1 - Initialize a grid full of walls

for ((i=0; i<maze_height*maze_width; i++)) ; do
    maze[$i]=1
done

# Step 2 - Pick a cell and add it to the maze. Add its neighboring walls to the wall list.

wall_list=()

maze[$(($maze_width+1))]=0

add_neighbors() {
    local cell_left=$(($1 - 1))
    local cell_right=$(($1 + 1))
    local cell_up=$(($1 - maze_width))
    local cell_down=$(($1 + maze_width))

    if ((cell_right % maze_width != maze_width - 1 && ${maze[cell_right]} == 1)) then
        wall_list[$cell_right]="r" 
    fi

    if ((cell_left % maze_width != 0 && ${maze[cell_left]} == 1)) then
        wall_list[$cell_left]="l"
    fi

    if ((cell_up > maze_width && ${maze[cell_up]} == 1)) then
        wall_list[$cell_up]="u"
    fi

    if ((cell_down < (maze_height-1)*maze_width && ${maze[cell_down]} == 1)) then
        wall_list[$cell_down]="d"
    fi
}

add_neighbors $(($maze_width+1))






print_maze() {
    clear
    for ((i = 0; i < maze_height; i++)) do
        for ((j = 0; j < maze_width; j++)) do
            if ((${maze[$((i * maze_width + j))]} == 1)) then
                echo -n ${wall_char}
            else
                echo -n " "
            fi
        done
        echo ""
    done
}

print_maze

echo ${wall_list[*]}