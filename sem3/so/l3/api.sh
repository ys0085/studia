let quote=$(curl "https://api.chucknorris.io/jokes/random" | jq -r '.value')
let imgurl=$(curl "https://api.thecatapi.com/v1/images/search" | jq -r '.[0].url')

curl ${imgurl} | catimg

echo ${quote}