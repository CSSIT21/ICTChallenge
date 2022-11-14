<script lang="ts">
	import type { Team } from 'src/types/leaderboard'
    import sit_logo from '../assets/images/sit-logo.png';
    import { fly } from 'svelte/transition';	
    import { flip } from 'svelte/animate'
    export let teams: Array<Team> = []
    let currentTime: Date
    let formattedTime: string
    setInterval(() => {  
        currentTime = new Date();
        var hr: number = currentTime.getHours()
        var timeSuffix = "AM"
        if (hr > 12) {
            hr -= 12
            timeSuffix = "PM"
        }
        var minute = currentTime.getMinutes()/10 > 1 ? currentTime.getMinutes() : "0" + currentTime.getMinutes()
        formattedTime = `${hr}:${minute} ${timeSuffix}`;

    }, 1000);
</script>

<main  class="bg-preview-bg w-screen h-screen relative bg-cover bg-no-repeat flex justify-center items-center flex-col" >
    <img class="absolute top-7 left-28" width="256px" src={sit_logo} alt="sit-logo">
    <h1 class="absolute top-12 right-24  text-zinc-50 text-4xl">{formattedTime}</h1>
    <h1 class="text-8xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-[#24DECC] to-[#FFFFFF]">SIT ICT CHALLENGE</h1>
    <div class="flex justify-evenly gap-40 flex-wrap p-44">
        {#each teams as team (team.id)}
            <div animate:flip={{ duration: (d) => 30 * Math.sqrt(d) }} transition:fly="{{ y: 200, duration: 2000 }}" class="animate-floating flex flex-col items-center" style="animation-delay: {Math.random()*1500}ms;">
                <h1 class="text-4xl font-semibold text-transparent bg-clip-text bg-gradient-to-b from-[#FFD303] to-[#E67E22]">{team.name}</h1>
                <h1 class="text-4xl text-zinc-50">{team.school}</h1>
            </div>        
        {/each}
    </div>
</main>
