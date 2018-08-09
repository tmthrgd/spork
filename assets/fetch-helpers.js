const headers = {
	Accept: 'text/plain',
};

const errorElem = document.querySelector('.error');

const fetchThen = resp => {
	if (resp.ok) {
		errorElem.textContent = '';
	} else {
		console.error('fetch request failed: ', resp);

		errorElem.textContent = `A HTTP error occurred: ${resp.status} ${resp.statusText}.`;
	}
};

const fetchCatch = e => {
	console.error('fetch request error: ', e);

	errorElem.textContent = `An error occurred: ${e.toString()}`;
};