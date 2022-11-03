<script lang="ts">
	import LeaderboardCard from "src/lib/components/LeaderboardCard.svelte"
	import Navbar from "src/lib/components/Navbar.svelte"
	import type { Team } from "src/types/leaderboard"

    let teams: Array<Team> = [
        {name:"muumel", order:1, school:"muumel", score:1000, isHighlighted:true},
        {name:"CYMPATI", order:2, school:"hello", score:500, isHighlighted:false},
        {name:"axxn", order:3, school:"hello", score:200, isHighlighted:false},
        {name:"muumel", order:4, school:"muumel", score:1000, isHighlighted:false},
        {name:"CYMPATI", order:5, school:"hello", score:500, isHighlighted:false},
        {name:"axxn", order:6, school:"hello", score:200, isHighlighted:false}
    ]

    function randomTeam() {
        clearInterval(random)
        var random = setInterval(() => {
            let random = Math.floor(Math.random() * teams.length)
            teams.forEach(element => {
                element.isHighlighted = false
            });
            teams[random].isHighlighted = true
        }, 250)
        setTimeout(() => {
            clearInterval(random)
        }, 5000);
    }
</script>

<main class="bg-[#1B2D51] h-screen w-screen px-24 py-12">
    <Navbar />
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <p on:click={randomTeam} class="text-white text-5xl font-semibold text-center mt-5">Leaderboard</p>
    {#each teams as team, i (i)}
        <LeaderboardCard name={team.name} order={team.order} score={team.score} isHighlighted={team.isHighlighted} />
    {/each}
</main>