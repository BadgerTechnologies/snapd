pipeline {
  agent {
    label "ros && small"
  }
  options {
    buildDiscarder(logRotator(daysToKeepStr: '360', numToKeepStr: '10',
                              artifactDaysToKeepStr: '360', artifactNumToKeepStr: '10'))
    timestamps()
    timeout(unit: 'MINUTES', time: 60)
  }
  stages {
    stage('prepare') {
      steps {
        sh script: """
          # The current setup of this snapcraft.yaml requires building with
          # 4.x. Specifically build-base being set to core is no longer allowed
          # with 5.x. When this is adjusted the version of snapcraft can be
          # allowed to be upgraded.
          sudo snap install --classic --channel=4.x/stable snapcraft
          sudo snap refresh snapcraft
          sudo snap install lxd
          export PATH=/snap/bin:$PATH
          sudo lxd.migrate
          lxd init --auto
        """
      }
    }
    stage('build') {
      steps {
        sh script: """
          export PATH=/snap/bin:$PATH
          snapcraft --use-lxd
        """
      }
      post {
        always {
          archiveArtifacts artifacts: "*.snap"
        }
      }
    }
    stage('s3-upload') {
      steps {
        withDockerRegistry(url: "https://242567060652.dkr.ecr.us-east-2.amazonaws.com", credentialsId:"ecr:us-east-2:CIUser") {
          withDockerContainer(image: '242567060652.dkr.ecr.us-east-2.amazonaws.com/ros-ci/awscli:v1.0') {
            withCredentials([[$class: 'AmazonWebServicesCredentialsBinding', credentialsId: 'CIUser']]) {
              sh script: """
                aws s3 cp snapd_*.snap s3://bar-builds/snapd/${BRANCH_NAME}/${GIT_COMMIT}/
              """
            }
          }
        }
      }
    }
  }
  post {
    always {
      cleanWs()
    }
  }
}
