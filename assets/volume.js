document.querySelector('.slider').addEventListener('input', e => {
	document.querySelector('.volume').textContent = e.target.value;

	fetch(`/volume/${encodeURIComponent(e.target.value)}`).catch(console.error);
}, false);