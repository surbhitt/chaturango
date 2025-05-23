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

function render_checkboard() {
    let container = document.getElementById('container')
    container.innerHTML = '';
    let board = document.createElement('table')
    board.innerHTML = '';

    for (let i = 0; i < 8; i++) {
        rank = document.createElement('tr')
        for (let j = 0; j < 8; j++) {
            file = document.createElement('td')
            // check if there is a piece 
            // put the piece
            if (board_state[i][j] != "") {
                console.log(board_state[i][j])
                let piece = document.createElement('div')
                piece.style.backgroundImage = "url(assets/pieces/" + board_state[i][j] + ".svg)" 
                piece.style.width = '100px'
                piece.style.height = '100px'
                // piece.style.background = 'red'
                piece.style.backgroundSize = 'cover'
                console.log("assets/pieces/" + board_state[i][j] + ".svg")
                file.appendChild(piece)
            }
            // check if it is selected
            if (i === selected_cell.x && j === selected_cell.y) {
                file.style.background = "red"
            }
            // assign x and y to data attributes
            file.dataset.x = i
            file.dataset.y = j
            // add event listener to check for click
            file.addEventListener('click', (event) => {
                selected_cell.x = Number(event.currentTarget.dataset.x)
                selected_cell.y = Number(event.currentTarget.dataset.y)
                render_checkboard()
            }
            )
            rank.appendChild(file)
        }
        board.appendChild(rank)
    }
    container.appendChild(board)
}

render_checkboard()
