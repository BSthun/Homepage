import { axios, caller } from './index';

export const trackLog = (event, beginningState, endingState) => {
	if (beginningState && typeof beginningState !== 'string') {
		beginningState = beginningState.toString();
	}
	if (endingState && typeof endingState !== 'string') {
		endingState = endingState.toString();
	}
	caller(axios.put(`/track/log/click`, {
		event: event, beginning_state: beginningState ?? '', ending_state: endingState ?? '',
	}));
};