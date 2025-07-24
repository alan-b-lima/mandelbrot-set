import "./wasm_exec.js";
const set = {
    origin: { x: 0, y: 0 },
    scale: -6,
    iterationLimit: 10
};
async function main() {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject);
    go.run(result.instance);
    render();
    window.addEventListener("keydown", onKeyDown);
}
function onKeyDown(evt) {
    switch (evt.key) {
        case "=":
            {
                set.scale--;
                render();
            }
            break;
        case "-":
            {
                set.scale++;
                render();
            }
            break;
        case "ArrowRight":
        case "d":
            {
                set.origin.x += 2 ** (set.scale + 3);
                render();
            }
            break;
        case "ArrowLeft":
        case "a":
            {
                set.origin.x -= 2 ** (set.scale + 3);
                render();
            }
            break;
        case "ArrowDown":
        case "s":
            {
                set.origin.y += 2 ** (set.scale + 3);
                render();
            }
            break;
        case "ArrowUp":
        case "w":
            {
                set.origin.y -= 2 ** (set.scale + 3);
                render();
            }
            break;
        case "i":
            {
                set.iterationLimit++;
                render();
            }
            break;
        case "I":
            {
                set.iterationLimit--;
                render();
            }
            break;
    }
}
function render() {
    const canvas = document.querySelector('canvas');
    const ctx = canvas.getContext('2d');
    const arr = generateMandelbrot(set.origin.x, set.origin.y, set.scale, canvas.width, canvas.height, set.iterationLimit * Math.max(1, -set.scale));
    const img = new ImageData(arr.buffer, arr.width, arr.height);
    ctx.putImageData(img, 0, 0, 0, 0, canvas.width, canvas.height);
}
window.addEventListener('DOMContentLoaded', main);
