let fib_sequence = [1,2];
let total = 2 

let index = 2

while (index <= 4000000){
    let value = fib_sequence[index-1] + fib_sequence[index-2]

    if (value > 4000000){
        break
    }
    fib_sequence.push(value)

    if( value % 2 == 0){
        total += value
    }
    

    index++
}

console.log(total)