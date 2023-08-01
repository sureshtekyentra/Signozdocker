/* eslint-disable @typescript-eslint/explicit-function-return-type */
/* eslint-disable object-shorthand */
/* eslint-disable func-names */

/**
 * Adds custom matchers from the react testing library to all tests
 */
import '@testing-library/jest-dom';
import 'jest-styled-components';

// Mock window.matchMedia
window.matchMedia =
	window.matchMedia ||
	function (): any {
		return {
			matches: false,
			addListener: function () {},
			removeListener: function () {},
		};
	};
