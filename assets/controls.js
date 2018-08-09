const headers = {
	Accept: 'text/plain',
};

document.querySelector('.controls').addEventListener('click', e => {
	if (!(e.target instanceof HTMLAnchorElement)) {
		return;
	}

	e.preventDefault();
	fetch(e.target.href, {headers}).catch(console.error);
}, false);

document.querySelector('.volume').addEventListener('input', e => {
	e.target.title = `${e.target.value}%`;

	fetch(`/controls/volume/${encodeURIComponent(e.target.value)}`, {headers}).catch(console.error);
}, false);