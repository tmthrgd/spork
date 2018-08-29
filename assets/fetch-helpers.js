const headers = {
	Accept: 'text/plain',
};

const errorElem = document.querySelector('.error');

const clearError = () => {
	errorElem.textContent = '';
};

const fetchThen = resp => {
	if (resp.ok) {
		clearError();
	} else {
		console.error('fetch request failed: ', resp);

		errorElem.textContent = `A HTTP error occurred: ${resp.status} ${resp.statusText}.`;
	}
};

const fetchCatch = e => {
	console.error('fetch request error: ', e);

	errorElem.textContent = `An error occurred: ${e.toString()}`;
};

const eventSourceError = e => {
	console.error('EventSource error: ', e);

	errorElem.textContent = 'The EventSource failed; unable to listen for playlist updates.';
};