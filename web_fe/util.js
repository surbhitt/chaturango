const { spawn } = require('child_process');
const readline = require('readline');

function xycoord_to_a1not(x, y) {
    // convert x, y coord to chess algebric notation
    let file = String.fromCharCode('a'.charCodeAt(0) + y)
    let rank = (x + 1).toString()

    return file + rank
}

function a1not_to_xycoord(notation) {
    return { x: Number(notation[1]) - 1, y: notation.charCodeAt(0) - 97 }
}

function start_engine() {
    const enginePath = '../bin/chaturam';
    const engine = spawn(enginePath);

    engine.stderr.on('data', (data) => {
        console.error(`Engine error: ${data}`);
    });

    return engine
}

function send_command(engine, command) {
    return new Promise((resolve) => {
        const rl = readline.createInterface({ input: engine.stdout });

        const responses = [];

        rl.on('line', (line) => {
            if (line.includes('bestmove') || line === 'readyok') {
                rl.close();
                resolve(responses);
            }
            else { responses.push(line) };
        });

        engine.stdin.write(command + '\n');
    });
}

module.exports = { start_engine, send_command, a1not_to_xycoord, xycoord_to_a1not };

//     // Handle stdout

//     // Handle stderr (optional, but useful for debugging)
//     engine.stderr.on('data', (data) => {
//         console.error(`Engine error: ${data.toString()}`);
//     });

//     // Handle exit
//     engine.on('close', (code) => {
//         console.log(`Engine exited with code ${code}`);
//     });

//     // Send a command to the engine (e.g., UCI protocol command)

//     // Example: Start with UCI handshake
//     sendCommand('uci');

//     // Later on, you can send more commands, for example:
//     setTimeout(() => {
//         sendCommand('isready');
//     }, 1000);

//     // To stop the engine gracefully
//     function stopEngine() {
//         sendCommand('quit');
//         engine.stdin.end();
//     }
