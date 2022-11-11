<script lang="ts">
	import type { Team } from "src/types/leaderboard"
        import PodiumComponent from "./PodiumComponent.svelte";
        export let teams: Array<Team>
</script>

<div class="relative h-[653px] flex justify-center gap-10 items-end">
        <div class="absolute bottom-0 w-screen scale-90">
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2  h-[256px] w-[512px] rounded-t-[384px]" />
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2 h-[384px] w-[768px] rounded-t-[576px]" />
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2 h-[512px] w-[1024px] rounded-t-[768px]" />
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2 h-[640px] w-[1280px] rounded-t-[960px]" />
        </div>
        
        {#each teams as team (team.id)}
                <PodiumComponent team={team.name} order={team.rank} score={team.score} />
        {/each}
        
</div>
