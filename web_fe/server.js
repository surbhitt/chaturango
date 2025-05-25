const express = require('express');
const util = require('./util')
const path = require('path');
const app = express();
const PORT = 3000;

app.use(express.json());
app.use(express.static(path.join(__dirname, 'public')));

const engine = util.start_engine()

app.post('/get-moves', async (req, res) => {
    const data = req.body;
    console.log('Received data:', data);
    let converted_notation = util.xycoord_to_a1not(data.x, data.y)
    console.log('converted notation:', converted_notation);
    let moves = await util.send_command(engine, converted_notation)
    console.log(moves)
    let converted_moves = []
    for (let i = 0; i < moves.length; i++) {
        converted_moves.push(util.a1not_to_xycoord(moves[i]))
    }
    console.log(converted_moves)
    res.json({
        moves: converted_moves
    });
});

app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
