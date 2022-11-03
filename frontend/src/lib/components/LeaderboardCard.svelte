<script lang="ts">
    import iconsCrystal from "../../assets/images/icons-crystal.png";
    import icons1st from "../../assets/images/icons-1stmedal-64.png";
    import icons2nd from "../../assets/images/icons-2ndmedal-64.png";
    import icons3rd from "../../assets/images/icons-3rdmedal-64.png";
	import { onMount } from "svelte"
    export let order: number
    export let name: string
    export let score: number
    export let isHighlighted: boolean

    let medal: string = ""
    onMount(() => {
        if (order == 1) {
            medal = icons1st
        } else if (order == 2) {
            medal = icons2nd
        } else if (order == 3) {
            medal = icons3rd
        }
    })
</script>
<!-- bg-gradient-to-r from-[#097FFE] to-[#8534E8] -->

<div class="flex gap-4 m-10 items-center">
    <p class="text-white text-5xl font-semibold">{order}</p>
    <div class="bg-opacity-20 w-full py-4 pl-16 pr-10 rounded-2xl shadow flex justify-between items-center relative transition-colors {isHighlighted ? "color" : "bg-white bg-opacity-20"}">
        <div class="text-4xl text-white">
            {#if medal !== ""}
                <img src={medal} class="absolute top-4 left-1 scale-75" alt="medal" />
            {/if}
            {name}
        </div>
        <div class="flex items-center text-4xl text-white w-40"><img src={iconsCrystal} alt="icons-crystal">{score}</div>
    </div>
</div>

<style>
    .color {
        background-image: linear-gradient(to right,#097FFE, #8534E8);
        background-color: rgb(255, 255, 255, 0.2);
        animation: changeColor 250ms;
    }
    @keyframes changeColor {
        0% {
            background-image: linear-gradient(to right,#097FFE, #8534E8);
            opacity: 0.7;
            background-color: rgb(255, 255, 255, 0.2);
        }
        100% {
            background-image: linear-gradient(to right,#097FFE, #8534E8);
            opacity: 1;
            background-color: rgb(255, 255, 255, 0.2);
        }
    }
</style>