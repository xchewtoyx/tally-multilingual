#!/usr/bin/perl

use strict;
use warnings;


sub main() {
    my %count;
    my @sorted;

    while(<>) {
        chomp;
        if( exists($count{$_}) ) {
            $count{$_}++;
        } else {
            $count{$_} = 1;
        }
    }

    @sorted = sort { $count{$b} <=> $count{$a} } keys(%count);

    for(my $i=0; $i<5; $i++) {
        my $line = sprintf("%4s %s\n", $sorted[$i], $count{$sorted[$i]});
        print $line;
    }
}

main unless caller;
