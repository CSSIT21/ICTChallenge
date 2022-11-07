<script lang="ts">
	import CheckCircleIcon from '../../assets/images/check_circle_FILL0_wght400_GRAD0_opsz48.svg'
	export let widthCard: string
	export let heightCard: string
	export let iconSize: string
	export let cardId: number = 0
	export let cardIndex: number = 0
	export let handleOpenModal: (id: number, index: number) => void = () => {}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div
	class="w-[{widthCard}] h-[{heightCard}] rounded-2xl bgColor flex justify-center items-center"
	on:click={() => {
		handleOpenModal(cardId, cardIndex)
	}}
>
	<img
		src={CheckCircleIcon}
		alt="icon"
		class="w-[{iconSize}] h-[{iconSize}]"
	/>
</div>

<style>
	.bgColor {
		background-color: rgba(255, 255, 255, 0.2);
	}
</style>
