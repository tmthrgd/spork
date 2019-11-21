import { headers, fetchThen, fetchCatch, playlistUpdates } from './helpers.js';

document.querySelector('.controls').addEventListener('click', e => {
	if (!(e.target instanceof HTMLAnchorElement)) {
		return;
	}

	let then = fetchThen;
	if ('toggle' in e.target.dataset) {
		then = resp => {
			fetchThen(resp);

			return resp.ok && resp.text().then(ok => {
				e.target.classList.toggle('active', ok.trim() === 'true');
			});
		};
	}

	e.preventDefault();
	fetch(e.target.href, { headers }).then(then).catch(fetchCatch);
}, false);

document.querySelector('.volume').addEventListener('input', e => {
	e.target.title = `${e.target.value}%`;

	fetch(`/controls/volume/${encodeURIComponent(e.target.value)}`, { headers })
		.then(fetchThen, fetchCatch);
}, false);

const song = document.querySelector('.song');

playlistUpdates(({ title, length }) => {
	if (title) {
		song.textContent = `${title} (${(length / 60) | 0}:${('0' + (length % 60)).slice(-2)})`;
	}
});