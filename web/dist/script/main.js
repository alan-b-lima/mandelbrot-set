import "./wasm_exec.js"

const set = {
    origin: { x: 0, y: 0 },
    scale: 0,
    iterationLimit: 20
}

async function main() {
    const go = new Go();

    const result = await WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject)
    go.run(result.instance)

    render()
    window.addEventListener("keydown", evt => {
        switch (evt.key) {

            case "=": {
                set.scale--
                render()
            } break

            case "-": {
                set.scale++
                render()
            } break

            case "d": {
                set.origin.x += 2 ** (set.scale + 3)
                render()
            } break

            case "a": {
                set.origin.x -= 2 ** (set.scale + 3)
                render()
            } break

            case "s": {
                set.origin.y += 2 ** (set.scale + 3)
                render()
            } break

            case "w": {
                set.origin.y -= 2 ** (set.scale + 3)
                render()
            } break

            case "i": {
                set.iterationLimit++
                render()
            } break

            case "I": {
                set.iterationLimit--
                render()
            } break
        }
    })
}

async function render() {
    const canvas = document.querySelector('canvas')
    const ctx = canvas.getContext('2d')

    const arr = generateMandelbrot(
        0, 0,
        2, 0,
        set.origin.x, set.origin.y,
        set.scale,
        canvas.width, canvas.height,
        set.iterationLimit)

    const img = new ImageData(arr, arr.width, arr.height)
    ctx.putImageData(img, 0, 0)
}

window.addEventListener('DOMContentLoaded', main)