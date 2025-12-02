const words = ["JABŁKO", "SAMOCHÓD", "KOMPUTER", "PROGRAM", "SZKOŁA", "PIŁKA", "MUZYKA", "ROWER"];
    const polishLetters = "AĄBCĆDEĘFGHIJKLŁMNŃOÓPRSŚTUVWXYZŹŻ".split("");

    let selectedWord = "";
    let guessedLetters = [];
    let wrongGuesses = 0;
    let gameActive = true;

    const wordContainer = document.getElementById("wordContainer");
    const lettersContainer = document.getElementById("lettersContainer");
    const canvas = document.getElementById("hangmanCanvas");
    const ctx = canvas.getContext("2d");
    const resultDiv = document.getElementById("result");

    function saveState() {
      const state = {
        selectedWord,
        guessedLetters,
        wrongGuesses,
        gameActive
      };
      localStorage.setItem("hangmanState", JSON.stringify(state));
    }

    function loadState() {
      const state = JSON.parse(localStorage.getItem("hangmanState"));
      if (state) {
        selectedWord = state.selectedWord;
        guessedLetters = state.guessedLetters;
        wrongGuesses = state.wrongGuesses;
        gameActive = state.gameActive;
        renderWord();
        renderLetters();
        drawHangman();
        checkWin();
        checkLose();
      } else {
        startGame(false);
      }
    }

    function startGame(newGame = false) {
      if (newGame || !localStorage.getItem("hangmanState")) {
        selectedWord = words[Math.floor(Math.random() * words.length)];
        guessedLetters = [];
        wrongGuesses = 0;
        gameActive = true;
        resultDiv.textContent = "";
      }
      drawHangman();
      renderWord();
      renderLetters();
      saveState();
    }

    function renderWord() {
      wordContainer.innerHTML = selectedWord.split("").map(letter => 
        guessedLetters.includes(letter) ? letter : "_"
      ).join(" ");
    }

    function renderLetters() {
      lettersContainer.innerHTML = "";
      polishLetters.forEach(letter => {
        const btn = document.createElement("div");
        btn.textContent = letter;
        btn.className = "letter";
        if (guessedLetters.includes(letter)) {
          btn.classList.add("used");
        }
        btn.addEventListener("click", () => guessLetter(letter));
        lettersContainer.appendChild(btn);
      });
    }

    function guessLetter(letter) {
      if (!gameActive) return;
      if (guessedLetters.includes(letter)) return;

      guessedLetters.push(letter);
      renderLetters();

      if (selectedWord.includes(letter)) {
        renderWord();
        checkWin();
      } else {
        wrongGuesses++;
        drawHangman();
        checkLose();
      }
      saveState();
    }

    function checkWin() {
      const won = selectedWord.split("").every(letter => guessedLetters.includes(letter));
      if (won) {
        gameActive = false;
        resultDiv.textContent = "Gratulacje! Wygrałeś!";
        saveState();
      }
    }

    function checkLose() {
      if (wrongGuesses >= 6) {
        gameActive = false;
        resultDiv.textContent = `Przegrałeś! Słowo to: ${selectedWord}`;
        saveState();
      }
    }

    function drawHangman() {
      ctx.clearRect(0, 0, canvas.width, canvas.height);
      // Podstawa
      ctx.beginPath(); ctx.moveTo(10, 190); ctx.lineTo(190, 190); ctx.stroke();
      // Słupek
      ctx.beginPath(); ctx.moveTo(50, 190); ctx.lineTo(50, 20); ctx.lineTo(120, 20); ctx.lineTo(120, 40); ctx.stroke();

      if (wrongGuesses > 0) { // Głowa
        ctx.beginPath(); ctx.arc(120, 55, 15, 0, Math.PI * 2); ctx.stroke();
      }
      if (wrongGuesses > 1) { // Tułów
        ctx.beginPath(); ctx.moveTo(120, 70); ctx.lineTo(120, 120); ctx.stroke();
      }
      if (wrongGuesses > 2) { // Lewa ręka
        ctx.beginPath(); ctx.moveTo(120, 80); ctx.lineTo(100, 100); ctx.stroke();
      }
      if (wrongGuesses > 3) { // Prawa ręka
        ctx.beginPath(); ctx.moveTo(120, 80); ctx.lineTo(140, 100); ctx.stroke();
      }
      if (wrongGuesses > 4) { // Lewa noga
        ctx.beginPath(); ctx.moveTo(120, 120); ctx.lineTo(100, 150); ctx.stroke();
      }
      if (wrongGuesses > 5) { // Prawa noga
        ctx.beginPath(); ctx.moveTo(120, 120); ctx.lineTo(140, 150); ctx.stroke();
      }
    }

    loadState();