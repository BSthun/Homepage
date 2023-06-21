<script lang="ts">
	import moment from 'moment'
	import { Link } from 'svelte-navigator'

	export let item: any
</script>

<div class="section-meta">
	{#if item.album != null}
		<Link to={'/photo/album/' + item.album.slug}>
			<div class="r-1 bordered">
				<i class="las la-arrow-left" />
				<div>
					<div class="r-1 subtitle">
						<i class="las la-images" />
						Album
					</div>
					<h4>{item.album.name}</h4>
				</div>
			</div>
		</Link>
	{/if}
	<h1>{item.section.title}</h1>
	<h2>{item.section.subtitle}</h2>
	<div class="row">
		<div class="r-1">
			<i class="las la-image" />
			<p>{item.section.photo_count}</p>
		</div>
		<div class="r-1">
			<i class="las la-clock" />
			<p>{moment(item.section.date).calendar()}</p>
		</div>
	</div>
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';

	.section-meta {
		display: flex;
		flex-direction: column;
		gap: 4px;

		h2 {
			font-size: 1rem;
		}
	}

	.row {
		display: flex;
		font-size: 1rem;

		@include breakpoint('md', 'dn') {
			flex-direction: column;
		}
	}

	.bordered {
		width: fit-content;
		line-height: 1.3;
		padding: 6px 12px;
		border-radius: 16px;
		border: 1px solid white;
		transition: background-color 0.2s ease-in-out;

		&:hover {
			background-color: rgba(255, 255, 255, 0.1);
		}
	}

	.subtitle {
		margin: 0;
		font-size: 0.7rem;
		font-weight: 100;

		i {
			font-size: 1em;
		}
	}
</style>
