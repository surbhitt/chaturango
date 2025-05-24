let selected_cell = { x: -1, y: -1 }

board_state = [
    ["R", "N", "B", "Q", "K", "B", "N", "R"],
    ["P", "P", "P", "P", "P", "P", "P", "P"],
    ["", "", "", "", "", "", "", ""],
    ["", "", "", "", "", "", "", ""],
    ["", "", "", "", "", "", "", ""],
    ["", "", "", "", "", "", "", ""],
    ["p", "p", "p", "p", "p", "p", "p", "p"],
    ["r", "n", "b", "q", "k", "b", "n", "r"]
]

function clear_highlighted_cells() {
    let highlighted_cells = document.getElementsByClassName('highlighted')
    for (let hc of highlighted_cells) {
        hc.classList.remove('highlighted')
    }
}

function highlight_moves(moves) {
    clear_highlighted_cells()

    for (let i = 0; i < moves.length; i++) {
        const x = moves[i].x;
        const y = moves[i].y;
        const cell = document.querySelector(`[data-x="${x}"][data-y="${y}"]`);
        if (cell) {
            cell.classList.add('highlighted')
        }
    }
}

async function get_moves() {
    // POST request to server
    // server then queries the engine
    try {
        let res = await fetch('http://localhost:3000/get-moves',
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(
                    { x: selected_cell.x, y: selected_cell.y }
                )
            }
        )
        res = await res.json()
        return res.moves
    } catch (err) {
        console.log("something went wrong")
        console.log(err)
        return []
    }
}

function render_checkboard() {
    let container = document.getElementById('container')
    container.innerHTML = '';
    let board = document.createElement('table')
    board.innerHTML = '';

    for (let i = 0; i < 8; i++) {
        rank = document.createElement('tr')
        for (let j = 0; j < 8; j++) {
            file = document.createElement('td')

            if (board_state[i][j] != "") {
                let piece = document.createElement('div')
                piece.style.backgroundImage = "url(assets/pieces/" + board_state[i][j] + ".svg)"
                piece.style.width = '100px'
                piece.style.height = '100px'
                piece.style.backgroundSize = 'cover'
                file.appendChild(piece)
            }

            file.dataset.x = i
            file.dataset.y = j

            file.addEventListener('click', async (element) => {
                // unselected selected cell
                if (selected_cell.x != -1) {
                    let cells = document.getElementsByClassName('selected')
                    cells[0].classList.remove('selected')
                }

                let piece = element.currentTarget
                piece.classList.add('selected')

                selected_cell.x = Number(element.currentTarget.dataset.x)
                selected_cell.y = Number(element.currentTarget.dataset.y)

                let moves = await get_moves()

                highlight_moves(moves)
            }
            )
            rank.appendChild(file)
        }
        board.appendChild(rank)
    }
    container.appendChild(board)
}

render_checkboard()
