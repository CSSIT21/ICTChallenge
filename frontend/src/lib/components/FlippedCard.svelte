<script lang="ts">
	import CheckCircleIcon from '../../assets/images/check_circle_FILL0_wght400_GRAD0_opsz48.svg'
	export let widthCard: string
	export let heightCard: string
	export let iconSize: string
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div
	class="rounded-2xl bgColor flex justify-center items-center"
	style="width: {widthCard}; height: {heightCard}"
>
	<img
		src={CheckCircleIcon}
		alt="icon"
		style="width: {iconSize}; height: {iconSize}"
	/>
</div>

<style>
	.bgColor {
		background-color: rgba(255, 255, 255, 0.2);
	}
</style>
