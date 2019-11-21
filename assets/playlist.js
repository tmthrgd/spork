import { headers, fetchThen, fetchCatch, playlistUpdates } from './helpers.js';

const playlist = document.querySelector('.playlist');

playlist.addEventListener('click', e => {
	if (!(e.target instanceof HTMLAnchorElement)) {
		return;
	}

	e.preventDefault();
	fetch(e.target.href, { headers }).then(fetchThen, fetchCatch);
}, false);

const items = playlist.getElementsByTagName('li');

playlistUpdates(({ pos }) => {
	const old = document.querySelector('.playlist .active');
	if (old) {
		old.id = '';
		old.classList.remove('active');
	}

	const item = items[pos];
	if (!item) {
		return;
	}

	item.id = 'current';
	item.classList.add('active');

	if (location.hash === '#current') {
		const scrollTo = item.previousElementSibling || item;
		scrollTo.scrollIntoView({
			block: 'start',
			inline: 'start',
			behavior: 'smooth',
		});
	}
});