// THIS FILE IS AUTO-GENERATED. You can regenerate it by running `go generate ./...` at the root of the lazygit repo.

package tests

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/bisect"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/branch"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/cherry_pick"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/commit"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/config"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/conflicts"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/custom_commands"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/diff"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/file"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/filter_and_search"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/filter_by_path"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/interactive_rebase"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/misc"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/patch_building"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/reflog"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/staging"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/stash"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/submodule"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/sync"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/tag"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/ui"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/undo"
)

var tests = []*components.IntegrationTest{
	bisect.Basic,
	bisect.FromOtherBranch,
	branch.CheckoutByName,
	branch.CreateTag,
	branch.Delete,
	branch.DetachedHead,
	branch.OpenWithCliArg,
	branch.Rebase,
	branch.RebaseAbortOnConflict,
	branch.RebaseAndDrop,
	branch.RebaseCancelOnConflict,
	branch.RebaseDoesNotAutosquash,
	branch.Reset,
	branch.ResetUpstream,
	branch.SetUpstream,
	branch.Suggestions,
	cherry_pick.CherryPick,
	cherry_pick.CherryPickConflicts,
	commit.Amend,
	commit.Commit,
	commit.CommitMultiline,
	commit.CommitWipWithPrefix,
	commit.CommitWithPrefix,
	commit.CreateTag,
	commit.DiscardOldFileChange,
	commit.Highlight,
	commit.History,
	commit.HistoryComplex,
	commit.NewBranch,
	commit.ResetAuthor,
	commit.Revert,
	commit.RevertMerge,
	commit.Reword,
	commit.Search,
	commit.SetAuthor,
	commit.StageRangeOfLines,
	commit.Staged,
	commit.StagedWithoutHooks,
	commit.Unstaged,
	config.RemoteNamedStar,
	conflicts.Filter,
	conflicts.ResolveExternally,
	conflicts.ResolveMultipleFiles,
	conflicts.UndoChooseHunk,
	custom_commands.BasicCmdAtRuntime,
	custom_commands.BasicCmdFromConfig,
	custom_commands.CheckForConflicts,
	custom_commands.ComplexCmdAtRuntime,
	custom_commands.FormPrompts,
	custom_commands.MenuFromCommand,
	custom_commands.MenuFromCommandsOutput,
	custom_commands.MultiplePrompts,
	custom_commands.OmitFromHistory,
	custom_commands.SuggestionsCommand,
	custom_commands.SuggestionsPreset,
	diff.Diff,
	diff.DiffAndApplyPatch,
	diff.DiffCommits,
	diff.IgnoreWhitespace,
	file.DirWithUntrackedFile,
	file.DiscardAllDirChanges,
	file.DiscardChanges,
	file.DiscardStagedChanges,
	file.DiscardUnstagedDirChanges,
	file.DiscardUnstagedFileChanges,
	file.Gitignore,
	file.RememberCommitMessageAfterFail,
	filter_and_search.FilterCommitFiles,
	filter_and_search.FilterFiles,
	filter_and_search.FilterMenu,
	filter_and_search.FilterRemoteBranches,
	filter_and_search.NestedFilter,
	filter_and_search.NestedFilterTransient,
	filter_by_path.CliArg,
	filter_by_path.SelectFile,
	filter_by_path.TypeFile,
	interactive_rebase.AdvancedInteractiveRebase,
	interactive_rebase.AmendCommitWithConflict,
	interactive_rebase.AmendFirstCommit,
	interactive_rebase.AmendFixupCommit,
	interactive_rebase.AmendHeadCommitDuringRebase,
	interactive_rebase.AmendMerge,
	interactive_rebase.AmendNonHeadCommitDuringRebase,
	interactive_rebase.DropTodoCommitWithUpdateRef,
	interactive_rebase.DropTodoCommitWithUpdateRefShowBranchHeads,
	interactive_rebase.DropWithCustomCommentChar,
	interactive_rebase.EditFirstCommit,
	interactive_rebase.EditNonTodoCommitDuringRebase,
	interactive_rebase.EditTheConflCommit,
	interactive_rebase.FixupFirstCommit,
	interactive_rebase.FixupSecondCommit,
	interactive_rebase.Move,
	interactive_rebase.MoveInRebase,
	interactive_rebase.MoveWithCustomCommentChar,
	interactive_rebase.PickRescheduled,
	interactive_rebase.Rebase,
	interactive_rebase.RewordCommitWithEditorAndFail,
	interactive_rebase.RewordFirstCommit,
	interactive_rebase.RewordLastCommit,
	interactive_rebase.RewordYouAreHereCommit,
	interactive_rebase.RewordYouAreHereCommitWithEditor,
	interactive_rebase.SquashDownFirstCommit,
	interactive_rebase.SquashDownSecondCommit,
	interactive_rebase.SquashFixupsAboveFirstCommit,
	interactive_rebase.SwapInRebaseWithConflict,
	interactive_rebase.SwapInRebaseWithConflictAndEdit,
	interactive_rebase.SwapWithConflict,
	misc.ConfirmOnQuit,
	misc.InitialOpen,
	misc.RecentReposOnLaunch,
	patch_building.Apply,
	patch_building.ApplyInReverse,
	patch_building.ApplyInReverseWithConflict,
	patch_building.CopyPatchToClipboard,
	patch_building.MoveToEarlierCommit,
	patch_building.MoveToEarlierCommitNoKeepEmpty,
	patch_building.MoveToIndex,
	patch_building.MoveToIndexPartOfAdjacentAddedLines,
	patch_building.MoveToIndexPartial,
	patch_building.MoveToIndexWithConflict,
	patch_building.MoveToLaterCommit,
	patch_building.MoveToLaterCommitPartialHunk,
	patch_building.MoveToNewCommit,
	patch_building.MoveToNewCommitPartialHunk,
	patch_building.RemoveFromCommit,
	patch_building.ResetWithEscape,
	patch_building.SelectAllFiles,
	patch_building.SpecificSelection,
	patch_building.StartNewPatch,
	reflog.Checkout,
	reflog.CherryPick,
	reflog.Patch,
	reflog.Reset,
	staging.DiffContextChange,
	staging.DiscardAllChanges,
	staging.Search,
	staging.StageHunks,
	staging.StageLines,
	staging.StageRanges,
	stash.Apply,
	stash.ApplyPatch,
	stash.CreateBranch,
	stash.Drop,
	stash.Pop,
	stash.PreventDiscardingFileChanges,
	stash.Rename,
	stash.Stash,
	stash.StashAll,
	stash.StashAndKeepIndex,
	stash.StashIncludingUntrackedFiles,
	stash.StashStaged,
	stash.StashUnstaged,
	submodule.Add,
	submodule.Enter,
	submodule.Remove,
	submodule.Reset,
	sync.FetchPrune,
	sync.ForcePush,
	sync.ForcePushMultipleMatching,
	sync.ForcePushMultipleUpstream,
	sync.Pull,
	sync.PullAndSetUpstream,
	sync.PullMerge,
	sync.PullMergeConflict,
	sync.PullRebase,
	sync.PullRebaseConflict,
	sync.PullRebaseInteractiveConflict,
	sync.PullRebaseInteractiveConflictDrop,
	sync.Push,
	sync.PushAndAutoSetUpstream,
	sync.PushAndSetUpstream,
	sync.PushFollowTags,
	sync.PushNoFollowTags,
	sync.PushTag,
	sync.PushWithCredentialPrompt,
	sync.RenameBranchAndPull,
	tag.Checkout,
	tag.CrudAnnotated,
	tag.CrudLightweight,
	tag.Reset,
	ui.Accordion,
	ui.DoublePopup,
	ui.EmptyMenu,
	ui.SwitchTabFromMenu,
	undo.UndoCheckoutAndDrop,
	undo.UndoDrop,
}
