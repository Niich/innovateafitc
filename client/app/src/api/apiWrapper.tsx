/**
 * Builds and sends a request to the server.
 *
 * errors can be shown to users by passing a notification hook
 *
 * @param input
 * @param param1
 */
export const apiWrapper = async (input: RequestInfo, { headers, ...init }: RequestInit): Promise<Response> => {

	return fetch(
		`${process.env.REACT_APP_API_ENDPOINT}${input}`,
		{ ...init, headers: { ...headers, 'Access-Control-Allow-Origin': 'http://localhost' } },
	).then((response: Response) => {
		if (!response.ok) {
			response.text().then((txt: string) => {
				console.log(txt)
			})
		}
		return response
	}).catch((err) => {
		console.log('Unable to communicate with endpoint. (っ °Д °;)っ')
		return err
	})
}
