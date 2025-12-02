

document.addEventListener("DOMContentLoaded", function() {
    const loadButton = document.getElementById("load-gallery");
    const galleryContainer = document.querySelector(".gallery-container");
    const loadingIndicator = document.querySelector(".loading-indicator");
    const galleryImages = [
        {
            alt: "3x3x3",
            src: "../images/3x3.jpg",
            title: "Kostka Rubika 3x3x3"
        },
        {
            alt: "7x7x7",
            src: "../images/7x7.jpg",
            title: "Kostka Rubika 7x7x7"
        },
        {
            alt: "Megaminx",
            src: "../images/megaminx.jpg",
            title: "Megaminx"
        },
        {
            alt: "Pyraminx",
            src: "../images/pyraminx.jpg",
            title: "Pyraminx"
        }
    ];
    function loadImage(imageInfo) {
        return new Promise(function (resolve, reject) {
            const img = new Image();
            img.onload = function () {
                const imgContainer = document.createElement("div");
                imgContainer.className = "gallery-item";
                const imgElement = document.createElement("img");
                imgElement.src = imageInfo.src;
                imgElement.alt = imageInfo.alt;
                imgElement.title = imageInfo.title;
                imgContainer.appendChild(imgElement);
                const caption = document.createElement("p");
                caption.textContent = imageInfo.title;
                imgContainer.appendChild(caption);
                resolve(imgContainer);
            };
            img.onerror = function () {
                reject(new Error(`Nie udało się załadować obrazu: ${imageInfo.src}`));
            };
            img.src = imageInfo.src;
        });
    }

    function loadGallery() {
        galleryContainer.innerHTML = "";
        loadingIndicator.style.display = "block";
        const imagePromises = galleryImages.map((imageInfo) => loadImage(imageInfo));
        Promise.allSettled(imagePromises)
            .then(function (imgContainers) {
                loadingIndicator.style.display = "none";
                imgContainers.forEach(function (container) {
                    if(container.status !== "rejected")
                        galleryContainer.appendChild(container.value);
                });
            });
    }

    loadGallery();
});