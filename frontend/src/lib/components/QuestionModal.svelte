<script lang="ts">
	import closeIcon from '../../assets/images/icons-close.svg'
	import bonusBg from '../../assets/images/bonus-question.png'
	import type { OpenQuestion } from 'src/types/question'

	export let open: boolean
	export let cardCol: number
	export let cardIndex: number
	export let openQuestion: OpenQuestion
	export let cardOpened: boolean
	export let handleCloseModal = (
		cardCol: number,
		cardIndex: number,
		opened: boolean
	) => {}

	let isCardOpen = true
	let timer: number = 15
	let isTimerRunning: boolean = false
	let isTimerFinish: boolean = false
	let minute: number = timer - 1
	let sec: number = 60

	const countdown = () => {
		sec--
		if (sec < 10) {
			document.getElementById('timer').innerHTML = minute + ' : 0' + sec
		} else if (sec == 60) {
			document.getElementById('timer').innerHTML = minute + 1 + ' : 00'
		} else {
			document.getElementById('timer').innerHTML = minute + ' : ' + sec
		}

		if (sec == 0) {
			minute--
			sec = 60
		}
		if (minute == -1 && sec == 60) {
			clearInterval(timerInterval)
			isTimerRunning = false
			isTimerFinish = true
		}
	}
	let timerInterval: any

	const countdownTimer = function (func: string) {
		if (func == 'start') {
			timerInterval = setInterval(countdown, 1000)
			isTimerRunning = true
		} else if (func == 'stop') {
			isTimerRunning = false
			clearInterval(timerInterval)
		}
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

{#if open}
	<div
		class="fixed top-0 left-0 z-50 w-full h-full bg-[#000000] {isCardOpen
			? 'opacity-50'
			: 'opacity-0'} modal-overlay transition-opacity duration-1000"
	/>
	<div
		class="z-50 w-[1600px] h-[872px] {isCardOpen
			? 'modal-container-slide-in'
			: 'modal-container-slide-out'} ml-20 flex flex-col mx-auto absolute overflow-y-auto question-bg-color rounded-[36px] shadow-xl items-center"
	>
		{#if openQuestion.bonus}
			<img src={bonusBg} alt="" class="absolute" />
		{/if}
		<div
			class="px-8 py-5 text-2xl font-extraboldhead absolute top-0 right-0"
		>
			<button
				class="p-2 ml-4 rounded-full cursor-pointer"
				on:click={() => {
					countdownTimer('stop')
					isCardOpen = false
					setTimeout(() => {
						handleCloseModal(cardCol, cardIndex, cardOpened)
						isCardOpen = true
					}, 900)
					timer = 15
					sec = 60
				}}
			>
				<img src={closeIcon} alt="closeIcon" class="w-12 h-12" />
			</button>
		</div>
		<div
			class="h-[678px] px-[200px] pt-12 text-center leading-[96px] flex justify-center items-center"
		>
			<p
				class="text-white {openQuestion.question.length < 85
					? 'text-[60px]'
					: 'text-[48px]'} font-semibold "
			>
				{openQuestion.question}
			</p>
		</div>
		<div class="flex justify-center ">
			<div
				class="h-[96px] w-[420px] flex flex-row justify-between items-center text-white font-medium"
			>
				<p
					class="text-2xl p-3 rounded-[10px] cursor-pointer
						{!isTimerRunning ? 'opacity-50' : 'opacity-100'} timer-bg"
					on:click={() =>
						isTimerRunning ? countdownTimer('stop') : null}
				>
					Stop
				</p>
				<div class="flex justify-center">
					<div
						class="rounded-[20px] h-[90px] flex items-center timer-bg"
					>
						<p class="text-[48px] px-6 text-center" id="timer">
							{timer} : 00
						</p>
					</div>
				</div>
				<p
					class="text-2xl p-3 rounded-[10px] cursor-pointer {isTimerFinish
						? 'opacity-50'
						: isTimerRunning
						? 'opacity-50'
						: 'opacity-100'} timer-bg"
					on:click={() =>
						!isTimerRunning && !isTimerFinish
							? countdownTimer('start')
							: null}
				>
					Start
				</p>
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
	.modal-container-slide-out {
		animation: slideOut 1s cubic-bezier(0.165, 0.84, 0.44, 1);
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
	@keyframes slideOut {
		0% {
			transform: translateY(0%);
		}
		100% {
			transform: translateY(110%);
		}
	}
</style>
