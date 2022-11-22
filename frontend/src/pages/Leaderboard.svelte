<script lang="ts">
	import LeaderboardCard from 'src/lib/components/LeaderboardCard.svelte'
	import PodiumBoard from 'src/lib/components/PodiumBoard.svelte'
	import PodiumSection from 'src/lib/components/PodiumSection.svelte'
	import { ArtWS } from 'src/store/websocket'
	import type { Team } from 'src/types/leaderboard'
	import { onDestroy } from 'svelte'
	import { flip } from 'svelte/animate'
	import Preview from './Preview.svelte'

	let mode = 'preview'
	let teams: Array<Team> = []
	let previewTeams: Array<Team> = []

	let teamsPodium: Array<Team> = []
    let teamsBoard: Array<Team> = []
	const client = ArtWS.connect(
	'ws://ictc-int.sit.kmutt.ac.th:3000/ws/projector/leaderboard?token=wdvXuDOytfx84J8d',
		{
			log: true, // Log for console.warning
			reconnect: true, // Reconnect on close
			reconnectInterval: 1000, // Reconnect interval in ms
			name: 'client', // Name for the client
		}
	)	

	const unsubscribeclient1 = client.subscribe('lb/state', (payload) => {
		mode = 'leaderboard'
		if (teams.length===0) {
			teams = payload.rankings.map((team:Team, i:number) => {
				return {...team, isHighlighted: false, rank: i + 1}
			})
		}else {
			console.log("reordering");
			updateScore(payload.rankings)
		}
		if (payload.highlighted_id) {
			resetHighlighted()
			setTimeout(() => {
				randomTeam(payload.highlighted_id)
			}, 3000);
		}
	})


	const unsubscribeclient2 = client.subscribe('lb/podium', (payload) => {
		mode = 'podium'
		teams = payload.rankings.map((team:Team, i:number) => {
			return {...team, isHighlighted: false, rank: i + 1}
		})
		startPodium()
	})

	const unsubscribeclient3 = client.subscribe('lb/preview', (payload) => {
		mode = 'preview'
	})

	const unsubscribeclient4 = client.subscribe('lb/preview/add', (payload) => {
		mode = 'preview'
		previewTeams = [...previewTeams, payload.team]
	})

	const unsubscribeclient5 = client.subscribe('lb/rankings', (payload) => {
		mode = 'leaderboard'
		if (teams.length===0) {
			teams = payload.rankings.map((team:Team, i:number) => {
				return {...team, isHighlighted: false, rank: i + 1}
			})
		}else {
			console.log("reordering");
			updateScore(payload.rankings)
		}
	})

	const unsubscribeclient6 = client.subscribe('lb/highlighted', (payload) => {
		mode = 'leaderboard'
		resetHighlighted()
		randomTeam(payload.highlighted_id)
	})

	onDestroy(() => {
		unsubscribeclient1()
		unsubscribeclient2()
		unsubscribeclient3()
		unsubscribeclient4()
		unsubscribeclient5()
		unsubscribeclient6()
	})

	function startPodium() {
		teamsPodium = [teams[1], teams[0], teams[2]];
        var insert = setInterval(()=>{
            if (teamsBoard.length === teams.length-4) {
                clearInterval(insert)
            }
            teamsBoard = [ ...teamsBoard,teams[3+teamsBoard.length]]
        },1000)
	}

	function randomTeam(id: number) {
		var random = setInterval(() => {
			let random = Math.floor(Math.random() * teams.length)
			resetHighlighted()
			teams[random].isHighlighted = true
		}, 250)
	
		setTimeout(() => {
			clearInterval(random)
			resetHighlighted()
			
			teams.find((el)=>el.id===id).isHighlighted = true
		}, 5000)

	}

	function resetHighlighted() {
		teams = teams.map((team) => {
			return { ...team, isHighlighted: false }
		})
	}

	function reorderDecend() {
		teams = teams.sort((a, b) => {
			if (a.score > b.score) {
				return -1
			} else if (b.score > a.score) {
				return 1
			} else {
				return 0
			}
		})
	}

	function updateScore(newTeams: Array<Team>) {
		teams = teams.map((team)=>{
			console.log(team.name,newTeams.find((t:Team)=>t.id === team.id).score);
			
			return {...team, score: newTeams.find((t:Team)=>t.id === team.id).score }
		})
		reorderDecend()
	}

</script>

	{#if mode === 'leaderboard'}
	<main class="bg-[#1B2D51] h-screen w-screen px-24 py-12">
		<p
		class="text-white text-5xl font-semibold text-center mt-5"
		>
			Leaderboard
		</p>
		{#each teams as team, i (team.rank)}
			<div animate:flip={{ duration: (d) => 30 * Math.sqrt(d) }}>
				<LeaderboardCard
					name={team.name}
					order={i + 1}
					score={team.score}
					isHighlighted={team.isHighlighted}
				/>
			</div>
		{/each}

	</main>
	{:else if mode === 'podium'}
	<main class="h-screen w-screen overflow-hidden bg-gradient-to-b from-[#3DC3B6] via-[#4F68BF] to-[#1B2D51]">
		<PodiumSection teams={teamsPodium} />
		<div class="p-2 mx-24 bg-[rgb(255,255,255,0.3)] rounded-t-3xl h-[421px]">
			<div class="py-10 px-32 flex flex-col gap-10 h-full">
				{#each teamsBoard as team (team.id)}
					<PodiumBoard name={team.name} order={team.rank} score={team.score} />
				{/each}
			</div>
		</div>
	</main>
	{:else}
	 <Preview teams={previewTeams} />
	{/if}
	
