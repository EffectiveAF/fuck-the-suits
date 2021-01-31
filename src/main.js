import 'string.prototype.startswith';
import './utils/object_values_polyfill.js';
import './utils/origin_polyfill.js';
import './utils/detect_mobile.js';

import App from './App.svelte';

const app = new App({
  target: document.body,
});

export default app;
