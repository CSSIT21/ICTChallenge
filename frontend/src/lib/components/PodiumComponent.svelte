<script lang="ts">
    import iconsCrystal from "../../assets/images/icons-crystal.png";
    import iconsMedieval from "../../assets/images/icons-medieval-crown.png";
    import icons2nd from "../../assets/images/icons-2ndmedal-64.png";
    import icons3rd from "../../assets/images/icons-3rdmedal-64.png";
    const partyJson = "src/assets/images/30172-flex-confetti.json"
    import { Lottie } from 'lottie-svelte';
	import { onMount } from "svelte"
    import { tweened } from 'svelte/motion';
    export let order: number;
    export let score: number;
    export let team: string;

    let gradient: string;
    let height: number;
    let width: number = 244;
    let color: string;
    let rankingIcon: string;
    let init: boolean = false;
    let showLabel: boolean = false;
    const tweenedScore = tweened(0, {duration: 2000})

    onMount(()=>{
        switch (order) {
            case 1:
                gradient = "from-[#FFEAC0] to-[#FFCC5B]";
                color = "from-[#FFA602] to-[#FFD33E]";
                height = 464;
                width = 292;
                rankingIcon = iconsMedieval;
                break;
            case 2:
                gradient = "from-[#E8EBEB] to-[#C1C7C8]";
                color = "from-[#9AA9AA] to-[#BABBBB]";
                height = 340;
                rankingIcon = icons2nd;
                break;
            case 3:
                gradient = "from-[#F9E2CD] to-[#ECB076]"
                color = "from-[#E67E22] to-[#E3A762]"
                height = 280;
                rankingIcon = icons3rd;
                break;
            default:
                break;
        }

        setTimeout(() => {            
            init = true;            
            setTimeout(() => {
                showLabel = true;
                tweenedScore.set(score)
            }, 3000); // 3000
        }, 3500*Math.abs(order-4)); // 3500
    })

</script>

<div style="height: {height}px; width: {width}px" class="relative {!showLabel&&"overflow-hidden block"}">
    {#if (showLabel&&order===1)}
    <div class="absolute z-[1000] scale-150 -top-40">
        <Lottie path={partyJson} speed={0.6}/>
    </div>
    {/if}
    
    <div id={"podium"+order}  class="{init ? "podium":"h-0"} relative flex flex-col items-center gap-7 bg-gradient-to-b {color} rounded-t-3xl">
        <div class="{!showLabel ? "hidden":"text-intro"}  absolute w-full -top-20 flex items-center justify-center text-white  text-4xl {order===1 && "flex-col font-semibold -top-28"}"><img class="{order!==1 && "scale-75"} {(showLabel&&order===1) &&  "winner"}" src={rankingIcon} alt="rank-icon"><div class="z-50 text-ellipsis overflow-hidden max-w-[120%] whitespace-nowrap">{team}</div></div>
        
        <div class="{!showLabel ? "hidden":"text-intro"} relative w-full mt-12 {order === 1 ? "h-[140px]": "h-[100px]"}">
            <div class="absolute left-1/2 -translate-x-1/2 leading-none font-bold text-transparent"  style="text-shadow: 0px 4.5478px 4.5478px rgba(0, 0, 0, 0.25); {order === 1 ? "font-size: 140px;": "font-size: 100px;"}">{order}</div>
            <div class="absolute left-1/2 -translate-x-1/2 leading-none font-bold text-transparent bg-clip-text bg-gradient-to-b {gradient}" style="{order === 1 ? "font-size: 140px;": "font-size: 100px;"}">{order}</div>
        </div>
        
        <div class="{!showLabel ? "hidden":"text-intro"} flex justify-center items-center rounded-2xl bg-[rgb(255,255,255,0.2)] px-4">
            <img src={iconsCrystal} alt="icons-crystal" class="scale-[60%]" />
            <div class="text-white text-3xl -translate-x-3">{Math.round($tweenedScore)}</div>
        </div>
    </div>
</div>
<style lang="scss">
    .podium {
        height: 100%;
        animation: increment 2s ease-in-out;
    }

    .winner {
        transition: 0.25s;
        animation: wiggling 2.5s ease-in-out;
    }

    .text-intro {
        animation: scaleup 1.5s ease-in-out;
    }

    @keyframes scaleup {
        0% {
            transform: scale(1);
        }
        50% {
            transform: scale(1.3);
        }
        100% {
            transform: scale(1);
        }
    }

    @keyframes increment {
        0% {
            transform: translateY(100%);
        }
        100%{
            transform: translateY(0);
        }
    }

    @keyframes wiggling {
        0%{
            transform: scale(1.2) rotate(-10deg) translateX(-50%);
        }
        50% {
            transform: scale(1.2) rotate(10deg) translateX(50%);
        }
        100% {
            transform: scale(1) rotate(0deg) translateX(0);
        }
    }
</style>