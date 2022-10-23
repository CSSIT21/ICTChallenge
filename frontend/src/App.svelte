<script lang="ts">
	import { Route, Router } from 'svelte-navigator'
	import About from './pages/About.svelte'
	import Home from './pages/Home.svelte'
</script>

<Router>
	<Route path="/">
		<Home />
	</Route>
	<Route path="/about">
		<About />
	</Route>
</Router>
