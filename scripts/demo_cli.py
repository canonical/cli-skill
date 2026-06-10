#!/usr/bin/env python3
"""Intentionally non-compliant demo CLI for task-management review tests."""

import argparse


def build_parser() -> argparse.ArgumentParser:
	parser = argparse.ArgumentParser(
		prog="taskThing",
		description="A wildly inconsistent task manager CLI that is intentionally bad for standards testing.",
		epilog="Try basically anything. Some commands probably look right.",
	)
	parser.add_argument("-v", "--verbose", action="store_true", help="be extra talkative")
	parser.add_argument("-q", "--quiet", action="store_true", help="say less maybe")
	parser.add_argument("--format", "-f", choices=["txt", "json", "table"], default="txt")

	commands = parser.add_subparsers(dest="what")

	add = commands.add_parser(
		"addTodo",
		help="Add a todo thing into the todo area",
		description="Create one todo item in a very flexible and not especially predictable way.",
	)
	add.add_argument("title", help="todo name")
	add.add_argument("place", nargs="?", help="where it lives")
	add.add_argument("owner", nargs="?", help="person maybe")
	add.add_argument("-p", "--priority", dest="priority")
	add.add_argument("--prio", dest="priority")
	add.add_argument("-t", "--tag", action="append")
	add.add_argument("--tags", help="comma separated tags")

	remove = commands.add_parser(
		"remove-todo-item",
		help="Remove, delete, or otherwise get rid of a todo",
		description="Deletes a task by whichever identifier feels useful at the time.",
	)
	remove.add_argument("todo")
	remove.add_argument("extra", nargs="?", help="second positional just because")
	remove.add_argument("-i", "--id")
	remove.add_argument("--name")
	remove.add_argument("-y", "--yes", action="store_true")
	remove.add_argument("--force", "-F", action="store_true")

	update = commands.add_parser(
		"update",
		help="Update something about todos",
		description="A two-level command for changing todo-ish data in a not very disciplined way.",
	)
	update_commands = update.add_subparsers(dest="update_target")
	update_todos = update_commands.add_parser(
		"todos",
		help="Update todo records in bulk or singular form",
	)
	update_todos.add_argument("todo")
	update_todos.add_argument("field", nargs="?", help="what to change")
	update_todos.add_argument("value", nargs="?", help="new value maybe")
	update_todos.add_argument("-s", "--status")
	update_todos.add_argument("--set-status")
	update_todos.add_argument("--label")
	update_todos.add_argument("--labels")

	read = commands.add_parser(
		"readTodos",
		help="Read todo information in one of several inconsistent ways",
		description="Shows one todo, many todos, or maybe all todos depending on the inputs you try.",
	)
	read.add_argument("todo", nargs="?", help="todo name or id or blank for everything")
	read.add_argument("-a", "--all", action="store_true")
	read.add_argument("--id")
	read.add_argument("--sort")
	read.add_argument("--order")
	read.add_argument("--no-headers", action="store_true")

	return parser


def main() -> None:
	build_parser().parse_args()


if __name__ == "__main__":
	main()
