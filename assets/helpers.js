export const headers = {
	Accept: 'text/plain',
};

const errorElem = document.querySelector('.error');

const clearError = () => {
	errorElem.textContent = '';
};

export function fetchThen(resp) {
	if (resp.ok) {
		clearError();
	} else {
		console.error('fetch request failed: ', resp);

		errorElem.textContent = `A HTTP error occurred: ${resp.status} ${resp.statusText}.`;
	}
};

export function fetchCatch(e) {
	console.error('fetch request error: ', e);

	errorElem.textContent = `An error occurred: ${e.toString()}`;
};

export function playlistUpdates(onmessage) {
	const es = new EventSource('/playlist/updates');

	es.onopen = clearError;

	es.onmessage = msg => {
		clearError();
		onmessage(JSON.parse(msg.data));
	};

	es.onerror = e => {
		console.error('EventSource error: ', e);

		errorElem.textContent = 'The EventSource failed; unable to listen for playlist updates.';
	};
}