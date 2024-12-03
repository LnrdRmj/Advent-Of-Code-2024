const fs = require('fs');
const readline = require('readline')

const fileStream = fs.createReadStream('input.txt');

(async function () {
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let cont = 0

    const list1 = []
    const list2 = []

    const mapNumberCounter = {}
    for await (const line of rl) {
        // Each line in input.txt will be successively available here as `line`.
        const [number1, number2] = line.split('   ').map(el => +el)
        list1.push(number1)
        list2.push(number2)

        mapNumberCounter[number2] = (mapNumberCounter[number2] ?? 0) + 1
    }

    list1.sort()
    list2.sort()

    // console.log(list1)
    // console.log(list2)

    // const result = list1.reduce((acc, el, index) => {
    //     return acc + Math.abs(list1[index] - list2[index])
    // }, 0)
    // console.log(result)

    const result = list1.reduce((acc, el, index) => {
        return acc + (mapNumberCounter[el] ?? 0) * el
    }, 0)
    console.log(result)

})()