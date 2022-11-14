<script lang="ts">
	import { Route, Router } from 'svelte-navigator'
	import About from './pages/About.svelte'
	import Home from './pages/Home.svelte'
	import Leaderboard from './pages/Leaderboard.svelte'
	import Preview from './pages/Preview.svelte'
	import Questions from './pages/Questions.svelte'
	import StudentPage from './pages/StudentPage.svelte'
	import Podium from './pages/Podium.svelte'
</script>

<Router>
	<Route path="/">
		<Preview />
	</Route>
	<Route path="/about">
		<About />
	</Route>
	<Route path="/podium">
		<Podium />
	</Route>
	<Route path="/leaderboard">
		<Leaderboard />
	</Route>
	<Route path="/question">
		<Questions />
	</Route>
	<Route path="/student">
		<StudentPage />
	</Route>
</Router>
