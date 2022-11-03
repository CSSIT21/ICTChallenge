<script lang="ts">
	import iconsCrystal from '../../assets/images/icons-crystal.png'

	export let img: string
	export let score: number
	export let textColor: string
	export let cardCol: number
	export let cardIndex: number
	export let openCard: (cardCol: number, cardIndex: number) => void
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div
	class="bg-white w-[324px] h-[180px] rounded-2xl flex flex-col justify-center items-center"
	on:click={() => openCard(cardCol, cardIndex)}
>
	<img src={img} alt="logo" class="w-[88px] mb-1" />
	<div class="flex items-center">
		<img src={iconsCrystal} alt="" class="w-[40px] h-[40px]" />
		<p class="text-[40px] font-semibold {textColor}">{score}</p>
	</div>
</div>

<style>
	.textColor1 {
		background: -webkit-linear-gradient(0deg, #0183ff, #00b2e3);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.textColor2 {
		background: -webkit-linear-gradient(0deg, #852ee6, #bc86f7);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.textColor3 {
		background: -webkit-linear-gradient(0deg, #ff9c1a, #ffd303);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.textColor4 {
		background: -webkit-linear-gradient(0deg, #f162ae, #ffc7e5);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>
