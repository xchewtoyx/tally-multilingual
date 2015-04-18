#!/usr/bin/php
<?php

function tally_lines($filename) {
    $counts = array();
    $handle = fopen($filename, "r");
    if ($handle) {
        while (($line = fgets($handle)) !== false) {
            $entry = trim($line);
            if( array_key_exists($entry, $counts) ) {
                $counts[$entry] += 1;
            } else {
                $counts[$entry] = 1;
            }
        }
        fclose($handle);
    }
    return $counts;
} 

function transform_counts($counts) {
    $combined = array();
    foreach($counts as $key => $value) {
        array_push($combined, array($key, $value));
    }
    return $combined;
}

function main() {
    global $argv;
    $counts = tally_lines($argv[1]);
    $combined = transform_counts($counts);
    usort($combined, function ($a, $b) {
        if( $a[1] == $b[1] ) {
            return 0;
        }
        return ($a[1]<$b[1])?1:-1;
    });
    for($i=0; $i<5; $i++) {
        printf("%4d %d\n", $combined[$i][0], $combined[$i][1]);
    }
}

main();
