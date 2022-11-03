<script lang="ts">
	import closeIcon from '../../assets/images/icons-close.svg'
	export let open: boolean
	export let cardCol: number
	export let cardIndex: number
	export let handleCloseModal = (cardCol: number, cardIndex: number) => {}
	let closeModal: boolean = false
</script>

{#if open}
	<div
		class="fixed top-0 left-0 z-50 flex items-center justify-center w-full h-full p-8 modal lg:p-0"
	>
		<div class="fixed w-full h-full bg-gray-900 opacity-50 modal-overlay" />
		<div
			class="z-50 w-[1600px] h-[872px] mx-auto overflow-y-auto question-bg-color rounded-[36px] shadow-xl modal-container"
		>
			<div
				class="flex items-center justify-end px-8 py-5 text-2xl font-extraboldhead"
			>
				<button
					class="p-2 ml-4 rounded-full hover:bg-gray-300"
					on:click={() => {
						handleCloseModal(cardCol, cardIndex)
					}}
				>
					<img src={closeIcon} alt="closeIcon" class="w-8 h-8" />
				</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.question-bg-color {
		background: -webkit-linear-gradient(
			0deg,
			rgba(92, 119, 171, 1),
			rgba(153, 164, 184, 1)
		);
	}
	.modal-container {
		animation: slideIn 0.5s linear;
	}
	.modal-slide-out {
		animation: slideOut 0.5s linear;
	}
	@keyframes slideIn {
		0% {
			transform: translateY(400px);
			animation-timing-function: ease-out;
		}
		60% {
			transform: translateY(-30px);
			animation-timing-function: ease-in;
		}
		80% {
			transform: translateY(10px);
			animation-timing-function: ease-out;
		}
		100% {
			transform: translateY(0px);
			animation-timing-function: ease-in;
		}
	}
	@keyframes slideOut {
		0% {
			opacity: 1;
		}
		100% {
			opacity: 0;
		}
	}
</style>
