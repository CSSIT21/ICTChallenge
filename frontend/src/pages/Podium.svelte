<script lang="ts">
	import PodiumBoard from "src/lib/components/PodiumBoard.svelte"
    import PodiumSection from "src/lib/components/PodiumSection.svelte";
	import { ArtWS } from "src/store/websocket"
	import type { Team } from "src/types/leaderboard"
	import { onMount } from "svelte"
    let teams: Array<Team> = []
    let teamsPodium: Array<Team> = []
    let teamsBoard: Array<Team> = []
    onMount(() => {
        teams = [
        {name:"Witty ICT women (อัจฉริยะICT)", school:"muumel", score:1000, isHighlighted:true, rank: 1, id: 1, percentile: 1},
        {name:"CYMPATI", school:"hello", score:500, isHighlighted:false, rank: 2, id: 2, percentile: 1},
        {name:"Witty ICT women (อัจฉริยะICT)", school:"hello", score:200, isHighlighted:false, rank: 3, id: 3, percentile: 1},
        {name:"muumel1", school:"muumel", score:1000, isHighlighted:false, rank: 4, id: 4, percentile: 1},
        {name:"CYMPATI1", school:"hello", score:500, isHighlighted:false,  rank: 5, id: 5, percentile: 1},
        {name:"axxn1", school:"hello", score:200, isHighlighted:false,  rank: 6, id: 6, percentile: 1}];
        teamsPodium = [teams[1], teams[0], teams[2]];
        var insert = setInterval(()=>{
            if (teamsBoard.length === teams.length-4) {
                clearInterval(insert)
            }
            teamsBoard = [ ...teamsBoard,teams[3+teamsBoard.length]]
            console.log(teamsBoard);
        },1000)
        
    })
    
</script>

<main class="h-screen w-screen overflow-hidden bg-gradient-to-b from-[#3DC3B6] via-[#4F68BF] to-[#1B2D51]">
    <PodiumSection teams={teamsPodium} />
    <div class="p-2 mx-24 bg-[rgb(255,255,255,0.3)] rounded-t-3xl h-[421px]">
        <div class="py-10 px-32 flex flex-col gap-10 h-full">
            {#each teamsBoard as team}
                <PodiumBoard name={team.name} order={team.rank} score={team.score} />
            {/each}
        </div>
    </div>
</main>

<style lang="scss">
    
</style>