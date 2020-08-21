import { apiWrapper } from './apiWrapper'
import { logEntry } from '../resources/interfaces'

export const getAllLogs = async (logEvent: Omit<logEntry, 'id'>): Promise<logEntry> => {
	return apiWrapper('/get-logs', {
		method: 'GET',
	}).then((response) => {
		if (response.ok) {
			return response.json()
		}
	})
}