/*
 * Copyright (c) 2023-present unTill Pro, Ltd.
 * @author: Victor Istratenko
 */
package main

func main() {

	cmdproc := buildCommandProcessor().
		addUpdateCmd().
		addDownloadCmd().
		addReleaseCmd().
		addGUICmd().
		addForkBranch().
		addDevBranch().
		addPr()
	cmdproc.Execute()
}
