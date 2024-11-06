import type { Reroute } from '@sveltejs/kit';
import { default as cookie } from 'js-cookie';

export const reroute: Reroute = ({ url }) => {
	return url.pathname;
};