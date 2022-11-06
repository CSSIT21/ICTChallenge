<script lang="ts">
	import { Route, Router } from 'svelte-navigator'
	import About from './pages/About.svelte'
	import Home from './pages/Home.svelte'
	import Leaderboard from './pages/Leaderboard.svelte'
	import Preview from './pages/Preview.svelte'
	import Podium from './pages/Podium.svelte'
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
	<Route path="/podium">
		<Podium />
	</Route>
	<Route path="/leaderboard">
		<Leaderboard />
	</Route>
</Router>
