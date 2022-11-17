<script lang="ts">
	import iconsCrystal from '../../assets/images/icons-crystal.png'

	export let img: string
	export let score: number
	export let cardId: number = 0
	export let topicId: number = 0
	export let textColor: string
	export let widthCard: string
	export let heightCard: string
	export let widthImg: string
	export let textSize: string
	export let handleFlipCard: (
		topicId: number,
		cardId: number
	) => void = () => {}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div>
	<div
		class="bg-white rounded-2xl flex flex-col justify-center items-center"
		style="width: {widthCard}; height: {heightCard}"
		on:click={() => {
			handleFlipCard(topicId, cardId)
		}}
	>
		<img
			src={img}
			alt="logo"
			class="mb-1"
			style="width: {widthImg} ; height: {widthImg}"
		/>
		<div class="flex items-center">
			<img
				src={iconsCrystal}
				alt=""
				style="width: {textSize}; height: {textSize}"
			/>
			<p
				class="font-semibold {textColor}"
				style="font-size: {textSize} ;"
			>
				{score}
			</p>
		</div>
	</div>
</div>

<style>
	.text-color-1 {
		background: -webkit-linear-gradient(0deg, #0183ff, #00b2e3);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-2 {
		background: -webkit-linear-gradient(0deg, #852ee6, #bc86f7);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-3 {
		background: -webkit-linear-gradient(0deg, #ff9c1a, #ffd303);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-4 {
		background: -webkit-linear-gradient(0deg, #f162ae, #ffc7e5);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-color-5 {
		background: -webkit-linear-gradient(0deg, #069a8e, #68d6cd);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>
