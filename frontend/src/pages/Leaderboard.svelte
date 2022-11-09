<script lang="ts">
	import LeaderboardCard from "src/lib/components/LeaderboardCard.svelte"
	import Navbar from "src/lib/components/Navbar.svelte"
	import type { Team } from "src/types/leaderboard"
	import { onMount } from "svelte"
    import { flip } from 'svelte/animate';

    let teams: Array<Team> = []
    onMount(() => {
        teams = [
        {name:"muumel", school:"muumel", score:1000, isHighlighted:true, rank: 1},
        {name:"CYMPATI", school:"hello", score:500, isHighlighted:false, rank: 2},
        {name:"axxn", school:"hello", score:200, isHighlighted:false, rank: 3},
        {name:"muumel1", school:"muumel", score:1000, isHighlighted:false, rank: 4},
        {name:"CYMPATI1", school:"hello", score:500, isHighlighted:false,  rank: 5},
        {name:"axxn1", school:"hello", score:200, isHighlighted:false,  rank: 6}]
    })
    

    function randomTeam() {
        var random = setInterval(() => {
            let random = Math.floor(Math.random() * teams.length);
            resetHighlighted();
            teams[random].isHighlighted = true;
        }, 250);
        setTimeout(() => {
            clearInterval(random)
        }, 5000);
    }

    function resetHighlighted() {
        teams = teams.map(team => {     
            return {...team,isHighlighted: false}
        }
        );        
    }

    function reorderDecend() {
        teams = teams.sort((a, b) => {
				if (a.score > b.score) {
					return -1;
				} else if (b.score > a.score) {
					return 1;
				} else {
					return 0;
				}
			});
    }

    function randomScore() {
		teams = teams.map((item) => {
			return {
				...item,
				score: Math.floor(Math.random() * 100)
			};
		});
        reorderDecend();
	}

	function resetScore() {
		teams = teams.map((item) => {
			return {
				...item,
				score: 0
			};
		});
		reorderDecend();
	}
</script>

<main class="bg-[#1B2D51] h-screen w-screen px-24 py-12">
    <Navbar />
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <p on:click={randomTeam} class="text-white text-5xl font-semibold text-center mt-5">Leaderboard</p>
    <button on:click={randomScore}>Rand Score</button>
    <button on:click={resetScore}>Reset Score</button>
    <button on:click={resetHighlighted}>Reset Highlighted</button>
    {#each teams as team, i (team.name)}
        <div animate:flip="{{duration:d => 30 * Math.sqrt(d)}}">
            <LeaderboardCard name={team.name} order={i+1} score={team.score} isHighlighted={team.isHighlighted} />
        </div>
    {/each}
</main>