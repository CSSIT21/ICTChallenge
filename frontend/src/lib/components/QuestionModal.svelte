<script lang="ts">
	import closeIcon from '../../assets/images/icons-close.svg'
	import bonusBg from '../../assets/images/bonus-question.png'
	import type { OpenQuestion } from 'src/types/question'
	import { onDestroy, onMount } from 'svelte'

	export let open: boolean
	export let openQuestion: OpenQuestion
	export let cardCol: number
	export let cardIndex: number
	export let handleCloseModal = (cardCol: number, cardIndex: number) => {}

	export let minute: number
	export let sec: number

	$: minute = minute
	$: sec = sec
	$: if (minute == 0 && sec == 0) {
		handleCloseModal(cardCol, cardIndex)
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

{#if open}
	<div
		class="fixed top-0 left-0 z-50 w-full h-full bg-[rgba(0,0,0,0.3)] opacity-50 modal-overlay transition-opacity duration-1000"
	/>
	<div
		class="z-50 w-[1600px] h-[972px] modal-container-slide-in ml-20 flex flex-col mx-auto absolute overflow-y-auto question-bg-color rounded-[36px] shadow-xl items-center"
	>
		{#if openQuestion.bonus}
			<img src={bonusBg} alt="" class="absolute -z-50" />
		{/if}
		<div
			class="h-[850px] px-[200px] text-center leading-[96px] flex flex-col justify-center items-center z-50"
		>
			<p
				class="text-white {openQuestion.question.title.length < 85
					? 'text-[60px]'
					: openQuestion.question.title.length > 120
					? 'text-[36px]'
					: 'text-[48px]'} font-semibold z-50"
			>
				{@html openQuestion.question.title}
			</p>

			{#if openQuestion.question.image_url != ''}
				<div class="w-[550px] flex justify-center">
					<img src={openQuestion.question.image_url} alt="img" />
				</div>
			{/if}
		</div>
		<div class="flex justify-center">
			<div
				class="h-[96px] flex flex-row justify-end items-end text-white font-medium"
			>
				<div class=" flex justify-center">
					<div
						class="rounded-[20px] h-[90px] flex items-center timer-bg"
					>
						<p class="text-[48px] px-6 text-center" id="timer">
							{minute} : {sec > 9 ? sec : '0' + sec}
						</p>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}

<style>
	.question-bg-color {
		background: linear-gradient(105.95deg, #445475, #2b3548);
	}
	.timer-bg {
		background: rgba(0, 0, 0, 0.1);
	}
	.modal-container-slide-in {
		animation: slideIn 1s linear;
	}

	@keyframes slideIn {
		0% {
			transform: translateY(100%);
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
</style>
