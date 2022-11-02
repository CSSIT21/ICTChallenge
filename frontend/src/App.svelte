<script lang="ts">
	import { Route, Router } from 'svelte-navigator'
	import About from './pages/About.svelte'
	import Home from './pages/Home.svelte'
	import Preview from './pages/Preview.svelte'
</script>

<Router>
	<Route path="/">
		<Home />
	</Route>
	<Route path="/about">
		<About />
	</Route>
	<Route path="/preview">
		<Preview />
	</Route>
</Router>
