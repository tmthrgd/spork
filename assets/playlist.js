const items = document.querySelectorAll('.playlist li');

const es = new EventSource('/playlist/updates');

es.onmessage = msg => {
	const {pos} = JSON.parse(msg.data);

	const old = document.querySelector('.playlist .active');
	if (old) {
		old.id = '';
		old.classList.remove('active');
	}

	const item = items[pos];
	if (item) {
		item.id = 'current';
		item.classList.add('active');
	}
};

es.onerror = e => {
	console.error(e);
};