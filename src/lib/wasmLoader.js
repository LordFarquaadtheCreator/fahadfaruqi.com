// @ts-ignore - Go is loaded from wasm_exec.js
export async function loadInterpreter() {
	// Load wasm_exec.js first
	await new Promise((resolve, reject) => {
		const script = document.createElement('script');
		script.src = '/wasm_exec.js';
		script.onload = resolve;
		script.onerror = reject;
		document.head.appendChild(script);
	});

	// @ts-ignore
	const go = new Go();

	const result = await WebAssembly.instantiateStreaming(
		fetch("/interp.wasm"),
		go.importObject
	);

	go.run(result.instance);

	// @ts-ignore - interpRun is exposed by Go WASM
	return (input) => window.interpRun(input);
}
