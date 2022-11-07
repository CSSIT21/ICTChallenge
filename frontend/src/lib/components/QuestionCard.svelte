<script lang="ts">
	import Modal from './QuestionModal.svelte'
	import iconsCrystal from '../../assets/images/icons-crystal.png'

	export let img: string
	export let score: number
	export let questionId: number
	export let textColor: string
	export let cardCol: number
	export let cardIndex: number
	export let openQuestion: string

	export let openedCard: (cardCol: number, cardIndex: number) => void
	export let getQuestion: (questionId: number) => void

	let showModal = false
	const handleOpenModal = () => {
		showModal = true
		getQuestion(questionId)
	}
	const handleCloseModal = (cardCol: number, cardIndex: number) => {
		showModal = false
		openedCard(cardCol, cardIndex)
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div
	class="bg-white w-[324px] h-[180px] rounded-2xl flex flex-col justify-center items-center cursor-pointer"
	on:click={() => {
		handleOpenModal()
	}}
>
	<img src={img} alt="logo" class="w-[88px] mb-1" />
	<div class="flex items-center">
		<img src={iconsCrystal} alt="" class="w-[40px] h-[40px]" />
		<p class="text-[40px] font-semibold {textColor}">{score}</p>
	</div>
</div>
<Modal
	open={showModal}
	{handleCloseModal}
	{cardCol}
	{cardIndex}
	{openQuestion}
/>

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
