import { axios, caller } from './index'

export const trackLog = (event, beginningState, endingState) => {
	caller(
		axios.put(`/track/log/click`, {
			event: event,
			beginning_state: beginningState,
			ending_state: endingState,
		})
	)
}
