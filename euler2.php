<?php 

$fib_sequence   = [1,2];
$total          = 2;
for($index = 2; $index <= 4_000_000; $index++){
    $value = $fib_sequence[$index-1] + $fib_sequence[$index - 2];
    if($value > 4_000_000){
        break;
    }
    
    $fib_sequence[] = $value;
    if($value % 2 == 0){
        $total += $value;
    }
}

echo "Total:{$total}\n";