<script lang="ts">
	import Modal from './QuestionModal.svelte'
	import iconsCrystal from '../../assets/images/icons-crystal.png'

	export let img: string
	export let score: number
	export let textColor: string
	export let cardCol: number
	export let cardIndex: number
	export let widthCard: string
	export let heightCard: string
	export let widthImg: string
	export let textSize: string
	export let openCard: (cardCol: number, cardIndex: number) => void
	let showModal = false
	const handleToggleModal = () => {
		showModal = !showModal
	}
	const handleCloseModal = (cardCol: number, cardIndex: number) => {
		showModal = false
		openCard(cardCol, cardIndex)
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div
	class="bg-white w-[{widthCard}] h-[{heightCard}] rounded-2xl flex flex-col justify-center items-center"
	on:click={() => {
		handleToggleModal()
	}}
>
	<img src={img} alt="logo" class="w-[{widthImg}] mb-1" />
	<div class="flex items-center">
		<img src={iconsCrystal} alt="" class="w-[{textSize}] h-[{textSize}]" />
		<p class="text-[{textSize}] font-semibold {textColor}">{score}</p>
	</div>
	<Modal open={showModal} {handleCloseModal} {cardCol} {cardIndex} />
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
</style>
