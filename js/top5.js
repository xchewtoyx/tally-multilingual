var fs = require('fs');
var readline = require('readline');
var stream = require('stream');
var util = require('util');

var filename = process.argv[2];

function tally(filename, callback) {
    var instream = fs.createReadStream(filename);
    var outstream = new stream;
    var rl = readline.createInterface(instream, outstream);

    var counter = {};

    rl.on('line', function(line) {
        var number = line.trim();
        if( counter[number] ) {
            counter[number] += 1;
        } else {
            counter[number] = 1;
        }
    });

    rl.on('close', function() { callback(counter) });
}

function top5(counter) {
    var key_list = Object.keys(counter);
    key_list.sort(function(a, b) {
        if( counter[a] == counter[b] ) {
            return 0;
        } else {
            return ( counter[a] > counter[b] )? -1: 1;
        }
    });
    return key_list.slice(0, 5);
}

function pad4(value) {
    return ('    ' + value).slice(-4);
}

tally(filename, function(counter) {
    var top_keys = top5(counter);
    for(var i=0; i<top_keys.length; i++) {
        var key = top_keys[i];
        console.log(pad4(key), counter[key]);
    }
});
