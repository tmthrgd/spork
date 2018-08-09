document.querySelector('.controls').addEventListener('click', e => {
	if (!(e.target instanceof HTMLAnchorElement)) {
		return;
	}

	e.preventDefault();
	fetch(e.target.href, {headers}).then(fetchThen, fetchCatch);
}, false);

document.querySelector('.volume').addEventListener('input', e => {
	e.target.title = `${e.target.value}%`;

	fetch(`/controls/volume/${encodeURIComponent(e.target.value)}`, {headers})
		.then(fetchThen, fetchCatch);
}, false);

const song = document.querySelector('.song');

const es = new EventSource('/playlist/updates');

es.onmessage = msg => {
	const {title, length} = JSON.parse(msg.data);
	if (title) {
		song.textContent = `${title} (${(length/60)|0}:${('0'+((length%60)|0)).slice(-2)})`;
	}
};

es.onerror = e => {
	console.error(e);
};