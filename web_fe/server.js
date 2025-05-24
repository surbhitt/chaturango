const express = require('express');
const path = require('path');
const app = express();
const PORT = 3000;

app.use(express.json());
app.use(express.static(path.join(__dirname, 'public')));

// app.get('/hello', (req, res) => {
//     res.send('Hello, world!');
// });

app.post('/get-moves', (req, res) => {
    const data = req.body;
    console.log('Received data:', data);
    res.json({
        moves: [
            {
                x: data.x + 1,
                y: data.y
            }
        ]
    });
});

app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
