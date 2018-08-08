document.querySelector('.volume').addEventListener('input', e => {
	e.target.title = `${e.target.value}%`;

	fetch(`/controls/volume/${encodeURIComponent(e.target.value)}`).catch(console.error);
}, false);